// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsDbFailedRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDrdsDbFailedRecordResponseBody
	GetRequestId() *string
	SetResult(v bool) *RemoveDrdsDbFailedRecordResponseBody
	GetResult() *bool
	SetSuccess(v bool) *RemoveDrdsDbFailedRecordResponseBody
	GetSuccess() *bool
}

type RemoveDrdsDbFailedRecordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D98BC610-5A91-453A-BC44-5873EF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database creation failure records were deleted from the DRDS instance.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDrdsDbFailedRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDrdsDbFailedRecordResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RemoveDrdsDbFailedRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetRequestId(v string) *RemoveDrdsDbFailedRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetResult(v bool) *RemoveDrdsDbFailedRecordResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) SetSuccess(v bool) *RemoveDrdsDbFailedRecordResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
