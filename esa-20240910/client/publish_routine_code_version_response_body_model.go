// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *PublishRoutineCodeVersionResponseBody
	GetCodeVersion() *string
	SetRequestId(v string) *PublishRoutineCodeVersionResponseBody
	GetRequestId() *string
}

type PublishRoutineCodeVersionResponseBody struct {
	// The code version.
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

func (s PublishRoutineCodeVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeVersionResponseBody) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *PublishRoutineCodeVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishRoutineCodeVersionResponseBody) SetCodeVersion(v string) *PublishRoutineCodeVersionResponseBody {
	s.CodeVersion = &v
	return s
}

func (s *PublishRoutineCodeVersionResponseBody) SetRequestId(v string) *PublishRoutineCodeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishRoutineCodeVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
