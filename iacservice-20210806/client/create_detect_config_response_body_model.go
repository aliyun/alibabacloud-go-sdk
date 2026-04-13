// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDetectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigId(v string) *CreateDetectConfigResponseBody
	GetDetectConfigId() *string
	SetRequestId(v string) *CreateDetectConfigResponseBody
	GetRequestId() *string
}

type CreateDetectConfigResponseBody struct {
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// example:
	//
	// String
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDetectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectConfigResponseBody) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *CreateDetectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDetectConfigResponseBody) SetDetectConfigId(v string) *CreateDetectConfigResponseBody {
	s.DetectConfigId = &v
	return s
}

func (s *CreateDetectConfigResponseBody) SetRequestId(v string) *CreateDetectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDetectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
