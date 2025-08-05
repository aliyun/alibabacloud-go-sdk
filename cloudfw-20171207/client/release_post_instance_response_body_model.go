// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePostInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpStatusCode(v int32) *ReleasePostInstanceResponseBody
	GetHttpStatusCode() *int32
	SetReleaseStatus(v bool) *ReleasePostInstanceResponseBody
	GetReleaseStatus() *bool
	SetRequestId(v string) *ReleasePostInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleasePostInstanceResponseBody
	GetSuccess() *bool
}

type ReleasePostInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Indicates whether the release was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ReleaseStatus *bool `json:"ReleaseStatus,omitempty" xml:"ReleaseStatus,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EE316A19-xxxx-5043-9DF1-C04458ABC570
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleasePostInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePostInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePostInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleasePostInstanceResponseBody) GetReleaseStatus() *bool {
	return s.ReleaseStatus
}

func (s *ReleasePostInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePostInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleasePostInstanceResponseBody) SetHttpStatusCode(v int32) *ReleasePostInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) SetReleaseStatus(v bool) *ReleasePostInstanceResponseBody {
	s.ReleaseStatus = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) SetRequestId(v string) *ReleasePostInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) SetSuccess(v bool) *ReleasePostInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
