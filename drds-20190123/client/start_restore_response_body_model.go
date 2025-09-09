// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRestoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRestoreResponseBody
	GetRequestId() *string
	SetResult(v string) *StartRestoreResponseBody
	GetResult() *string
	SetSuccess(v bool) *StartRestoreResponseBody
	GetSuccess() *bool
}

type StartRestoreResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 721C71DD-D3D0-4327-BFDD-678326******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SQL audit was disabled for the DRDS database.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartRestoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRestoreResponseBody) GoString() string {
	return s.String()
}

func (s *StartRestoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRestoreResponseBody) GetResult() *string {
	return s.Result
}

func (s *StartRestoreResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartRestoreResponseBody) SetRequestId(v string) *StartRestoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRestoreResponseBody) SetResult(v string) *StartRestoreResponseBody {
	s.Result = &v
	return s
}

func (s *StartRestoreResponseBody) SetSuccess(v bool) *StartRestoreResponseBody {
	s.Success = &v
	return s
}

func (s *StartRestoreResponseBody) Validate() error {
	return dara.Validate(s)
}
