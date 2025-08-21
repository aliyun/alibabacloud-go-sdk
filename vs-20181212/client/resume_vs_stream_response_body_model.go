// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVsStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeVsStreamResponseBody
	GetRequestId() *string
}

type ResumeVsStreamResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeVsStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeVsStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeVsStreamResponseBody) SetRequestId(v string) *ResumeVsStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeVsStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
