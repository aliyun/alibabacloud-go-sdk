// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackupsSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveBackupsSetResponseBody
	GetRequestId() *string
	SetResult(v string) *RemoveBackupsSetResponseBody
	GetResult() *string
	SetSuccess(v bool) *RemoveBackupsSetResponseBody
	GetSuccess() *bool
}

type RemoveBackupsSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 70FED5BE-4DDC-4556-AD35-5A6D27******
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

func (s RemoveBackupsSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackupsSetResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBackupsSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveBackupsSetResponseBody) GetResult() *string {
	return s.Result
}

func (s *RemoveBackupsSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveBackupsSetResponseBody) SetRequestId(v string) *RemoveBackupsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) SetResult(v string) *RemoveBackupsSetResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) SetSuccess(v bool) *RemoveBackupsSetResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveBackupsSetResponseBody) Validate() error {
	return dara.Validate(s)
}
