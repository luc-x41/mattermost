// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// ReactionStore is an autogenerated mock type for the ReactionStore type
type ReactionStore struct {
	mock.Mock
}

// BulkGetForPosts provides a mock function with given fields: postIds
func (_m *ReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, error) {
	ret := _m.Called(postIds)

	var r0 []*model.Reaction
	if rf, ok := ret.Get(0).(func([]string) []*model.Reaction); ok {
		r0 = rf(postIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(postIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: reaction
func (_m *ReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, error) {
	ret := _m.Called(reaction)

	var r0 *model.Reaction
	if rf, ok := ret.Get(0).(func(*model.Reaction) *model.Reaction); ok {
		r0 = rf(reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Reaction) error); ok {
		r1 = rf(reaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAllWithEmojiName provides a mock function with given fields: emojiName
func (_m *ReactionStore) DeleteAllWithEmojiName(emojiName string) error {
	ret := _m.Called(emojiName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(emojiName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedRowsByIds provides a mock function with given fields: r
func (_m *ReactionStore) DeleteOrphanedRowsByIds(r *model.RetentionIdsForDeletion) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.RetentionIdsForDeletion) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetForPost provides a mock function with given fields: postID, allowFromCache
func (_m *ReactionStore) GetForPost(postID string, allowFromCache bool) ([]*model.Reaction, error) {
	ret := _m.Called(postID, allowFromCache)

	var r0 []*model.Reaction
	if rf, ok := ret.Get(0).(func(string, bool) []*model.Reaction); ok {
		r0 = rf(postID, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(postID, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForPostSince provides a mock function with given fields: postId, since, excludeRemoteId, inclDeleted
func (_m *ReactionStore) GetForPostSince(postId string, since int64, excludeRemoteId string, inclDeleted bool) ([]*model.Reaction, error) {
	ret := _m.Called(postId, since, excludeRemoteId, inclDeleted)

	var r0 []*model.Reaction
	if rf, ok := ret.Get(0).(func(string, int64, string, bool) []*model.Reaction); ok {
		r0 = rf(postId, since, excludeRemoteId, inclDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, string, bool) error); ok {
		r1 = rf(postId, since, excludeRemoteId, inclDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopForTeamSince provides a mock function with given fields: teamID, userID, since, offset, limit
func (_m *ReactionStore) GetTopForTeamSince(teamID string, userID string, since int64, offset int, limit int) (*model.TopReactionList, error) {
	ret := _m.Called(teamID, userID, since, offset, limit)

	var r0 *model.TopReactionList
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) *model.TopReactionList); ok {
		r0 = rf(teamID, userID, since, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TopReactionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int64, int, int) error); ok {
		r1 = rf(teamID, userID, since, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopForUserSince provides a mock function with given fields: userID, teamID, since, offset, limit
func (_m *ReactionStore) GetTopForUserSince(userID string, teamID string, since int64, offset int, limit int) (*model.TopReactionList, error) {
	ret := _m.Called(userID, teamID, since, offset, limit)

	var r0 *model.TopReactionList
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) *model.TopReactionList); ok {
		r0 = rf(userID, teamID, since, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TopReactionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int64, int, int) error); ok {
		r1 = rf(userID, teamID, since, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *ReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	ret := _m.Called(endTime, limit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(endTime, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteByUser provides a mock function with given fields: userID
func (_m *ReactionStore) PermanentDeleteByUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: reaction
func (_m *ReactionStore) Save(reaction *model.Reaction) (*model.Reaction, error) {
	ret := _m.Called(reaction)

	var r0 *model.Reaction
	if rf, ok := ret.Get(0).(func(*model.Reaction) *model.Reaction); ok {
		r0 = rf(reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Reaction) error); ok {
		r1 = rf(reaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
