// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackupLocalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetBackupLocalResponseBody
	GetRequestId() *string
	SetResult(v string) *SetBackupLocalResponseBody
	GetResult() *string
	SetSuccess(v bool) *SetBackupLocalResponseBody
	GetSuccess() *bool
}

type SetBackupLocalResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6F70CE62-5077-4B7B-95BC-4DAC45614DBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
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

func (s SetBackupLocalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackupLocalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetBackupLocalResponseBody) GetResult() *string {
	return s.Result
}

func (s *SetBackupLocalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetBackupLocalResponseBody) SetRequestId(v string) *SetBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackupLocalResponseBody) SetResult(v string) *SetBackupLocalResponseBody {
	s.Result = &v
	return s
}

func (s *SetBackupLocalResponseBody) SetSuccess(v bool) *SetBackupLocalResponseBody {
	s.Success = &v
	return s
}

func (s *SetBackupLocalResponseBody) Validate() error {
	return dara.Validate(s)
}
