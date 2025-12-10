// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentsUsersStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetExperimentsUsersStatisticsResponseBody
	GetRequestId() *string
	SetUsers(v []*GetExperimentsUsersStatisticsResponseBodyUsers) *GetExperimentsUsersStatisticsResponseBody
	GetUsers() []*GetExperimentsUsersStatisticsResponseBodyUsers
}

type GetExperimentsUsersStatisticsResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*GetExperimentsUsersStatisticsResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s GetExperimentsUsersStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsUsersStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentsUsersStatisticsResponseBody) GetUsers() []*GetExperimentsUsersStatisticsResponseBodyUsers {
	return s.Users
}

func (s *GetExperimentsUsersStatisticsResponseBody) SetRequestId(v string) *GetExperimentsUsersStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentsUsersStatisticsResponseBody) SetUsers(v []*GetExperimentsUsersStatisticsResponseBodyUsers) *GetExperimentsUsersStatisticsResponseBody {
	s.Users = v
	return s
}

func (s *GetExperimentsUsersStatisticsResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExperimentsUsersStatisticsResponseBodyUsers struct {
	// example:
	//
	// 12345******67890
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetExperimentsUsersStatisticsResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsUsersStatisticsResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *GetExperimentsUsersStatisticsResponseBodyUsers) SetUserId(v string) *GetExperimentsUsersStatisticsResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *GetExperimentsUsersStatisticsResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
