// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackDriftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriftDetectionId(v string) *DetectStackDriftResponseBody
	GetDriftDetectionId() *string
	SetRequestId(v string) *DetectStackDriftResponseBody
	GetRequestId() *string
}

type DetectStackDriftResponseBody struct {
	// The ID of the drift detection.
	//
	// example:
	//
	// a7044f0d-6f2e-4128-a307-4524ef88****
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectStackDriftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectStackDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackDriftResponseBody) GetDriftDetectionId() *string {
	return s.DriftDetectionId
}

func (s *DetectStackDriftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectStackDriftResponseBody) SetDriftDetectionId(v string) *DetectStackDriftResponseBody {
	s.DriftDetectionId = &v
	return s
}

func (s *DetectStackDriftResponseBody) SetRequestId(v string) *DetectStackDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackDriftResponseBody) Validate() error {
	return dara.Validate(s)
}
