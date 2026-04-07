// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZooKeeperSaslUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListZooKeeperSaslUserResponseBodyData) *ListZooKeeperSaslUserResponseBody
	GetData() []*ListZooKeeperSaslUserResponseBodyData
	SetRequestId(v string) *ListZooKeeperSaslUserResponseBody
	GetRequestId() *string
}

type ListZooKeeperSaslUserResponseBody struct {
	Data []*ListZooKeeperSaslUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 6C075654-E42F-5F58-914F-E11DA70881B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListZooKeeperSaslUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListZooKeeperSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListZooKeeperSaslUserResponseBody) GetData() []*ListZooKeeperSaslUserResponseBodyData {
	return s.Data
}

func (s *ListZooKeeperSaslUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListZooKeeperSaslUserResponseBody) SetData(v []*ListZooKeeperSaslUserResponseBodyData) *ListZooKeeperSaslUserResponseBody {
	s.Data = v
	return s
}

func (s *ListZooKeeperSaslUserResponseBody) SetRequestId(v string) *ListZooKeeperSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZooKeeperSaslUserResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListZooKeeperSaslUserResponseBodyData struct {
	// example:
	//
	// 1631001140913
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// UserAddNeedReload
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListZooKeeperSaslUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListZooKeeperSaslUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListZooKeeperSaslUserResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListZooKeeperSaslUserResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListZooKeeperSaslUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListZooKeeperSaslUserResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *ListZooKeeperSaslUserResponseBodyData) SetCreateTime(v string) *ListZooKeeperSaslUserResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListZooKeeperSaslUserResponseBodyData) SetDescription(v string) *ListZooKeeperSaslUserResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListZooKeeperSaslUserResponseBodyData) SetStatus(v string) *ListZooKeeperSaslUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListZooKeeperSaslUserResponseBodyData) SetUserName(v string) *ListZooKeeperSaslUserResponseBodyData {
	s.UserName = &v
	return s
}

func (s *ListZooKeeperSaslUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
