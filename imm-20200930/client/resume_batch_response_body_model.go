// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeBatchResponseBody
	GetRequestId() *string
}

type ResumeBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D2C628B8-35DF-473C-8A41-757F30******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeBatchResponseBody) SetRequestId(v string) *ResumeBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
