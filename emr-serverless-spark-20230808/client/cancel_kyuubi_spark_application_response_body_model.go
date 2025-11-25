// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelKyuubiSparkApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CancelKyuubiSparkApplicationResponseBodyBody) *CancelKyuubiSparkApplicationResponseBody
	GetBody() *CancelKyuubiSparkApplicationResponseBodyBody
	SetRequestId(v string) *CancelKyuubiSparkApplicationResponseBody
	GetRequestId() *string
}

type CancelKyuubiSparkApplicationResponseBody struct {
	Body *CancelKyuubiSparkApplicationResponseBodyBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelKyuubiSparkApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelKyuubiSparkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelKyuubiSparkApplicationResponseBody) GetBody() *CancelKyuubiSparkApplicationResponseBodyBody {
	return s.Body
}

func (s *CancelKyuubiSparkApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelKyuubiSparkApplicationResponseBody) SetBody(v *CancelKyuubiSparkApplicationResponseBodyBody) *CancelKyuubiSparkApplicationResponseBody {
	s.Body = v
	return s
}

func (s *CancelKyuubiSparkApplicationResponseBody) SetRequestId(v string) *CancelKyuubiSparkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelKyuubiSparkApplicationResponseBody) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelKyuubiSparkApplicationResponseBodyBody struct {
	// example:
	//
	// spark-339f844005b6404c95f9f7c7a13b****
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelKyuubiSparkApplicationResponseBodyBody) String() string {
	return dara.Prettify(s)
}

func (s CancelKyuubiSparkApplicationResponseBodyBody) GoString() string {
	return s.String()
}

func (s *CancelKyuubiSparkApplicationResponseBodyBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CancelKyuubiSparkApplicationResponseBodyBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelKyuubiSparkApplicationResponseBodyBody) SetApplicationId(v string) *CancelKyuubiSparkApplicationResponseBodyBody {
	s.ApplicationId = &v
	return s
}

func (s *CancelKyuubiSparkApplicationResponseBodyBody) SetSuccess(v bool) *CancelKyuubiSparkApplicationResponseBodyBody {
	s.Success = &v
	return s
}

func (s *CancelKyuubiSparkApplicationResponseBodyBody) Validate() error {
	return dara.Validate(s)
}
