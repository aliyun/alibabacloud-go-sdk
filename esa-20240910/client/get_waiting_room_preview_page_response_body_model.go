// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWaitingRoomPreviewPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageUrl(v string) *GetWaitingRoomPreviewPageResponseBody
	GetPageUrl() *string
	SetRequestId(v string) *GetWaitingRoomPreviewPageResponseBody
	GetRequestId() *string
}

type GetWaitingRoomPreviewPageResponseBody struct {
	// The waiting room preview page URL.
	//
	// example:
	//
	// http://waitingroom.myalicdn.com/testxxxx
	PageUrl *string `json:"PageUrl,omitempty" xml:"PageUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWaitingRoomPreviewPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWaitingRoomPreviewPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetWaitingRoomPreviewPageResponseBody) GetPageUrl() *string {
	return s.PageUrl
}

func (s *GetWaitingRoomPreviewPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWaitingRoomPreviewPageResponseBody) SetPageUrl(v string) *GetWaitingRoomPreviewPageResponseBody {
	s.PageUrl = &v
	return s
}

func (s *GetWaitingRoomPreviewPageResponseBody) SetRequestId(v string) *GetWaitingRoomPreviewPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWaitingRoomPreviewPageResponseBody) Validate() error {
	return dara.Validate(s)
}
