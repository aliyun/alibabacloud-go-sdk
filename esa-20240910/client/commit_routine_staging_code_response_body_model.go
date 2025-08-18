// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitRoutineStagingCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *CommitRoutineStagingCodeResponseBody
	GetCodeVersion() *string
	SetRequestId(v string) *CommitRoutineStagingCodeResponseBody
	GetRequestId() *string
}

type CommitRoutineStagingCodeResponseBody struct {
	// The version number of the newly generated code.
	//
	// example:
	//
	// 1710120201067203242
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitRoutineStagingCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommitRoutineStagingCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CommitRoutineStagingCodeResponseBody) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *CommitRoutineStagingCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommitRoutineStagingCodeResponseBody) SetCodeVersion(v string) *CommitRoutineStagingCodeResponseBody {
	s.CodeVersion = &v
	return s
}

func (s *CommitRoutineStagingCodeResponseBody) SetRequestId(v string) *CommitRoutineStagingCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitRoutineStagingCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
