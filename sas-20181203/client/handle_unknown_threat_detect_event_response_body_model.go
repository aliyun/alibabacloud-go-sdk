// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleUnknownThreatDetectEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *HandleUnknownThreatDetectEventResponseBody
	GetRequestId() *string
}

type HandleUnknownThreatDetectEventResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6EBB8614-746D-555D-AB69-C801AE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleUnknownThreatDetectEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleUnknownThreatDetectEventResponseBody) GoString() string {
	return s.String()
}

func (s *HandleUnknownThreatDetectEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleUnknownThreatDetectEventResponseBody) SetRequestId(v string) *HandleUnknownThreatDetectEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleUnknownThreatDetectEventResponseBody) Validate() error {
	return dara.Validate(s)
}
