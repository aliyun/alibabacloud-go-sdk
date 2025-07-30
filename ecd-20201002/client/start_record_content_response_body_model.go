// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRecordContentResponseBody
	GetRequestId() *string
}

type StartRecordContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRecordContentResponseBody) SetRequestId(v string) *StartRecordContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRecordContentResponseBody) Validate() error {
	return dara.Validate(s)
}
