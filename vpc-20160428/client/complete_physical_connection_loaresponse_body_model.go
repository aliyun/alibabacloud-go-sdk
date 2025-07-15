// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompletePhysicalConnectionLOAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompletePhysicalConnectionLOAResponseBody
	GetRequestId() *string
}

type CompletePhysicalConnectionLOAResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F8983C74-E068-4509-B442-89BD82C8F43B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompletePhysicalConnectionLOAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompletePhysicalConnectionLOAResponseBody) GoString() string {
	return s.String()
}

func (s *CompletePhysicalConnectionLOAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompletePhysicalConnectionLOAResponseBody) SetRequestId(v string) *CompletePhysicalConnectionLOAResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompletePhysicalConnectionLOAResponseBody) Validate() error {
	return dara.Validate(s)
}
