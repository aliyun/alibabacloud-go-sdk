// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateDetectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateDetectConfigResponseBody
	GetRequestId() *string
}

type DissociateDetectConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DissociateDetectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateDetectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateDetectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateDetectConfigResponseBody) SetRequestId(v string) *DissociateDetectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateDetectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
