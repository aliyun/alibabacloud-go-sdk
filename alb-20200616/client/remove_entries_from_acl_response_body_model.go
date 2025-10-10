// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntriesFromAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *RemoveEntriesFromAclResponseBody
	GetJobId() *string
	SetRequestId(v string) *RemoveEntriesFromAclResponseBody
	GetRequestId() *string
}

type RemoveEntriesFromAclResponseBody struct {
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
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntriesFromAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntriesFromAclResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *RemoveEntriesFromAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveEntriesFromAclResponseBody) SetJobId(v string) *RemoveEntriesFromAclResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) SetRequestId(v string) *RemoveEntriesFromAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) Validate() error {
	return dara.Validate(s)
}
