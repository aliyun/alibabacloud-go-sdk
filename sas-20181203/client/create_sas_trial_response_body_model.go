// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSasTrialResponseBody
	GetRequestId() *string
}

type CreateSasTrialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9D22BDB7-C0**328A2B2E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSasTrialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSasTrialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSasTrialResponseBody) SetRequestId(v string) *CreateSasTrialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSasTrialResponseBody) Validate() error {
	return dara.Validate(s)
}
