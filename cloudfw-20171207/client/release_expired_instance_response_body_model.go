// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseExpiredInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpStatusCode(v int32) *ReleaseExpiredInstanceResponseBody
	GetHttpStatusCode() *int32
	SetReleaseStatus(v string) *ReleaseExpiredInstanceResponseBody
	GetReleaseStatus() *string
	SetRequestId(v string) *ReleaseExpiredInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseExpiredInstanceResponseBody
	GetSuccess() *bool
}

type ReleaseExpiredInstanceResponseBody struct {
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// true
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" xml:"ReleaseStatus,omitempty"`
	// example:
	//
	// 9EC1DB0F-EE53-5D36-B5DA-71CB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseExpiredInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleaseExpiredInstanceResponseBody) GetReleaseStatus() *string {
	return s.ReleaseStatus
}

func (s *ReleaseExpiredInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseExpiredInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseExpiredInstanceResponseBody) SetHttpStatusCode(v int32) *ReleaseExpiredInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseExpiredInstanceResponseBody) SetReleaseStatus(v string) *ReleaseExpiredInstanceResponseBody {
	s.ReleaseStatus = &v
	return s
}

func (s *ReleaseExpiredInstanceResponseBody) SetRequestId(v string) *ReleaseExpiredInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseExpiredInstanceResponseBody) SetSuccess(v bool) *ReleaseExpiredInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseExpiredInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
