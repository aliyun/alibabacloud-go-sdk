// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteAclResponseBody
	GetJobId() *string
	SetRequestId(v string) *DeleteAclResponseBody
	GetRequestId() *string
}

type DeleteAclResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DeleteAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAclResponseBody) SetJobId(v string) *DeleteAclResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclResponseBody) Validate() error {
	return dara.Validate(s)
}
