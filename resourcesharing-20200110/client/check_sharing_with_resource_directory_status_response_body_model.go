// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSharingWithResourceDirectoryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableSharingWithRd(v bool) *CheckSharingWithResourceDirectoryStatusResponseBody
	GetEnableSharingWithRd() *bool
	SetRequestId(v string) *CheckSharingWithResourceDirectoryStatusResponseBody
	GetRequestId() *string
}

type CheckSharingWithResourceDirectoryStatusResponseBody struct {
	// Indicates whether resource sharing within a resource directory is enabled. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	EnableSharingWithRd *bool `json:"EnableSharingWithRd,omitempty" xml:"EnableSharingWithRd,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819545D0-C97A-5DB3-BD73-A1B17E9A4BC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSharingWithResourceDirectoryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSharingWithResourceDirectoryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) GetEnableSharingWithRd() *bool {
	return s.EnableSharingWithRd
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) SetEnableSharingWithRd(v bool) *CheckSharingWithResourceDirectoryStatusResponseBody {
	s.EnableSharingWithRd = &v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) SetRequestId(v string) *CheckSharingWithResourceDirectoryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSharingWithResourceDirectoryStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
