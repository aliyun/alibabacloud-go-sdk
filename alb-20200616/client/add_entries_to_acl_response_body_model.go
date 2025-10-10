// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntriesToAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *AddEntriesToAclResponseBody
	GetJobId() *string
	SetRequestId(v string) *AddEntriesToAclResponseBody
	GetRequestId() *string
}

type AddEntriesToAclResponseBody struct {
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
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEntriesToAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEntriesToAclResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *AddEntriesToAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEntriesToAclResponseBody) SetJobId(v string) *AddEntriesToAclResponseBody {
	s.JobId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) SetRequestId(v string) *AddEntriesToAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) Validate() error {
	return dara.Validate(s)
}
