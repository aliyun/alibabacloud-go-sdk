// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQuerySessionByClientIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineStatusList(v []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) *BatchQuerySessionByClientIdsResponseBody
	GetOnlineStatusList() []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList
	SetRequestId(v string) *BatchQuerySessionByClientIdsResponseBody
	GetRequestId() *string
}

type BatchQuerySessionByClientIdsResponseBody struct {
	// The list of client statuses.
	OnlineStatusList []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList `json:"OnlineStatusList,omitempty" xml:"OnlineStatusList,omitempty" type:"Repeated"`
	// The request ID. It is unique for each request.
	//
	// example:
	//
	// 63309FDB-ED6C-46AE-B31C-A172FBA0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchQuerySessionByClientIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponseBody) GetOnlineStatusList() []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	return s.OnlineStatusList
}

func (s *BatchQuerySessionByClientIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQuerySessionByClientIdsResponseBody) SetOnlineStatusList(v []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) *BatchQuerySessionByClientIdsResponseBody {
	s.OnlineStatusList = v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBody) SetRequestId(v string) *BatchQuerySessionByClientIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBody) Validate() error {
	if s.OnlineStatusList != nil {
		for _, item := range s.OnlineStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQuerySessionByClientIdsResponseBodyOnlineStatusList struct {
	// The client ID of the client. For more information about client IDs, see [Glossary](https://help.aliyun.com/document_detail/42420.html).
	//
	// example:
	//
	// GID_test@0001
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Indicates whether the device is online. Valid values:
	//
	// - **true**: Online.
	//
	// - **false**: Offline.
	//
	// example:
	//
	// true
	OnlineStatus *bool `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
}

func (s BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) String() string {
	return dara.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) GetClientId() *string {
	return s.ClientId
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) GetOnlineStatus() *bool {
	return s.OnlineStatus
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) SetClientId(v string) *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	s.ClientId = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) SetOnlineStatus(v bool) *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	s.OnlineStatus = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) Validate() error {
	return dara.Validate(s)
}
