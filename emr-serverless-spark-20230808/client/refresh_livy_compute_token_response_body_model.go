// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshLivyComputeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefreshLivyComputeTokenResponseBody
	GetCode() *string
	SetMessage(v string) *RefreshLivyComputeTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefreshLivyComputeTokenResponseBody
	GetRequestId() *string
}

type RefreshLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefreshLivyComputeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *RefreshLivyComputeTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefreshLivyComputeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshLivyComputeTokenResponseBody) SetCode(v string) *RefreshLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshLivyComputeTokenResponseBody) SetMessage(v string) *RefreshLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshLivyComputeTokenResponseBody) SetRequestId(v string) *RefreshLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshLivyComputeTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
