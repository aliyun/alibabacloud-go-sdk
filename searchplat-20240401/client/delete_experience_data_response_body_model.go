// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperienceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExperienceDataResponseBody
	GetRequestId() *string
}

type DeleteExperienceDataResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 65C620DA-D6BE-5F56-BBCD-6F2282BB7BAD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteExperienceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperienceDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperienceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperienceDataResponseBody) SetRequestId(v string) *DeleteExperienceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperienceDataResponseBody) Validate() error {
	return dara.Validate(s)
}
