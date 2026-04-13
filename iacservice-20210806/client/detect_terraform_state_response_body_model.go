// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectTerraformStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectionId(v string) *DetectTerraformStateResponseBody
	GetDetectionId() *string
	SetRequestId(v string) *DetectTerraformStateResponseBody
	GetRequestId() *string
}

type DetectTerraformStateResponseBody struct {
	// example:
	//
	// job-dcsdxxxxxx
	DetectionId *string `json:"detectionId,omitempty" xml:"detectionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BF75EF50-955D-5E1F-AB23-A657C2C6D3C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DetectTerraformStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectTerraformStateResponseBody) GoString() string {
	return s.String()
}

func (s *DetectTerraformStateResponseBody) GetDetectionId() *string {
	return s.DetectionId
}

func (s *DetectTerraformStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectTerraformStateResponseBody) SetDetectionId(v string) *DetectTerraformStateResponseBody {
	s.DetectionId = &v
	return s
}

func (s *DetectTerraformStateResponseBody) SetRequestId(v string) *DetectTerraformStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectTerraformStateResponseBody) Validate() error {
	return dara.Validate(s)
}
