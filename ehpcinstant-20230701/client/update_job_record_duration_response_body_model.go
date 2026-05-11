// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRecordDurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateJobRecordDurationResponseBody
	GetRequestId() *string
}

type UpdateJobRecordDurationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobRecordDurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRecordDurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobRecordDurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobRecordDurationResponseBody) SetRequestId(v string) *UpdateJobRecordDurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobRecordDurationResponseBody) Validate() error {
	return dara.Validate(s)
}
