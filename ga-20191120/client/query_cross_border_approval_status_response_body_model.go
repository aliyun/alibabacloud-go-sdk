// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCrossBorderApprovalStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *QueryCrossBorderApprovalStatusResponseBody
	GetApprovalStatus() *string
	SetRequestId(v string) *QueryCrossBorderApprovalStatusResponseBody
	GetRequestId() *string
}

type QueryCrossBorderApprovalStatusResponseBody struct {
	// Cross border permissions of Alibaba Cloud account (main account).
	//
	// -  UNAPPLIED : No cross-border permission application has been submitted or application records cannot be found.
	//
	// -  APPLIED : Cross-border permission review in progress.
	//
	// -  REJECTED : Cross-border permission review failed.
	//
	// -  PASSED : Cross-border permission review passed.
	//
	// example:
	//
	// UNAPPLIED
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F96E634B-A523-587F-9A09-AE8B2FD04B9C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCrossBorderApprovalStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCrossBorderApprovalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCrossBorderApprovalStatusResponseBody) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *QueryCrossBorderApprovalStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCrossBorderApprovalStatusResponseBody) SetApprovalStatus(v string) *QueryCrossBorderApprovalStatusResponseBody {
	s.ApprovalStatus = &v
	return s
}

func (s *QueryCrossBorderApprovalStatusResponseBody) SetRequestId(v string) *QueryCrossBorderApprovalStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCrossBorderApprovalStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
