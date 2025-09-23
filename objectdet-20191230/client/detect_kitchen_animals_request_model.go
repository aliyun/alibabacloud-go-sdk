// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectKitchenAnimalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLA(v string) *DetectKitchenAnimalsRequest
	GetImageURLA() *string
	SetImageURLB(v string) *DetectKitchenAnimalsRequest
	GetImageURLB() *string
}

type DetectKitchenAnimalsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectKitchenAnimals/DetectKitchenAnimals-left1.png
	ImageURLA *string `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectKitchenAnimals/DetectKitchenAnimals-right1.png
	ImageURLB *string `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
}

func (s DetectKitchenAnimalsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsRequest) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsRequest) GetImageURLA() *string {
	return s.ImageURLA
}

func (s *DetectKitchenAnimalsRequest) GetImageURLB() *string {
	return s.ImageURLB
}

func (s *DetectKitchenAnimalsRequest) SetImageURLA(v string) *DetectKitchenAnimalsRequest {
	s.ImageURLA = &v
	return s
}

func (s *DetectKitchenAnimalsRequest) SetImageURLB(v string) *DetectKitchenAnimalsRequest {
	s.ImageURLB = &v
	return s
}

func (s *DetectKitchenAnimalsRequest) Validate() error {
	return dara.Validate(s)
}
