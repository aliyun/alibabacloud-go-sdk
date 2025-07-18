// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeDataRedistributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeDataRedistributeResponseBody
	GetRequestId() *string
}

type ResumeDataRedistributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeDataRedistributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeDataRedistributeResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeDataRedistributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeDataRedistributeResponseBody) SetRequestId(v string) *ResumeDataRedistributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeDataRedistributeResponseBody) Validate() error {
	return dara.Validate(s)
}
