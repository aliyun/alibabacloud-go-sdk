// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomBlockRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomBlockRecordResponseBody
	GetRequestId() *string
}

type DeleteCustomBlockRecordResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomBlockRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomBlockRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomBlockRecordResponseBody) SetRequestId(v string) *DeleteCustomBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomBlockRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
