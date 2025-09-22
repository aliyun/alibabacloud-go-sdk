// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFolderListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ChatFolderItem) *GetChatFolderListResponseBody
	GetData() []*ChatFolderItem
	SetRequestId(v string) *GetChatFolderListResponseBody
	GetRequestId() *string
}

type GetChatFolderListResponseBody struct {
	// Returned data
	Data []*ChatFolderItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// A8AEC6D9-A359-5169-BD1A-BD848BA60D65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChatFolderListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatFolderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatFolderListResponseBody) GetData() []*ChatFolderItem {
	return s.Data
}

func (s *GetChatFolderListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatFolderListResponseBody) SetData(v []*ChatFolderItem) *GetChatFolderListResponseBody {
	s.Data = v
	return s
}

func (s *GetChatFolderListResponseBody) SetRequestId(v string) *GetChatFolderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatFolderListResponseBody) Validate() error {
	return dara.Validate(s)
}
