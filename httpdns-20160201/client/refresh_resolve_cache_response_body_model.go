// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshResolveCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefreshResolveCacheResponseBody
	GetCode() *string
	SetMessage(v string) *RefreshResolveCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefreshResolveCacheResponseBody
	GetRequestId() *string
}

type RefreshResolveCacheResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FA8C2599-362D-4113-8FB4-E88A40C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshResolveCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshResolveCacheResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshResolveCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *RefreshResolveCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefreshResolveCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshResolveCacheResponseBody) SetCode(v string) *RefreshResolveCacheResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshResolveCacheResponseBody) SetMessage(v string) *RefreshResolveCacheResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshResolveCacheResponseBody) SetRequestId(v string) *RefreshResolveCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshResolveCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
