// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultWatermarkResponseBody
	GetRequestId() *string
}

type SetDefaultWatermarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultWatermarkResponseBody) SetRequestId(v string) *SetDefaultWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}
