// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcknowledgeTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcknowledgeTaskResultResponseBody
	GetRequestId() *string
	SetResult(v int32) *AcknowledgeTaskResultResponseBody
	GetResult() *int32
}

type AcknowledgeTaskResultResponseBody struct {
	// example:
	//
	// D6CB3623-4726-4947-AC2B-2C6E673B447C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AcknowledgeTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcknowledgeTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *AcknowledgeTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcknowledgeTaskResultResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *AcknowledgeTaskResultResponseBody) SetRequestId(v string) *AcknowledgeTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcknowledgeTaskResultResponseBody) SetResult(v int32) *AcknowledgeTaskResultResponseBody {
	s.Result = &v
	return s
}

func (s *AcknowledgeTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
