// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetBackupPolicyResponseBody
	GetRequestId() *string
	SetResult(v string) *SetBackupPolicyResponseBody
	GetResult() *string
	SetSuccess(v bool) *SetBackupPolicyResponseBody
	GetSuccess() *bool
}

type SetBackupPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A3140FC7-B78B-4D8E-B0C8-926D28******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup policy was successfully configured.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the database creation failure records were removed from the DRDS instance.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetBackupPolicyResponseBody) GetResult() *string {
	return s.Result
}

func (s *SetBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetBackupPolicyResponseBody) SetRequestId(v string) *SetBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackupPolicyResponseBody) SetResult(v string) *SetBackupPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *SetBackupPolicyResponseBody) SetSuccess(v bool) *SetBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *SetBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
