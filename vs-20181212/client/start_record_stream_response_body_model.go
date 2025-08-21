// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRecordStreamResponseBody
	GetRequestId() *string
}

type StartRecordStreamResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRecordStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRecordStreamResponseBody) SetRequestId(v string) *StartRecordStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRecordStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
