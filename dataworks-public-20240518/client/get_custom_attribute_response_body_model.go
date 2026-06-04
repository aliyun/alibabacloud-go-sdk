// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAttribute(v *CustomAttribute) *GetCustomAttributeResponseBody
	GetCustomAttribute() *CustomAttribute
	SetRequestId(v string) *GetCustomAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomAttributeResponseBody
	GetSuccess() *bool
}

type GetCustomAttributeResponseBody struct {
	CustomAttribute *CustomAttribute `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BA51C9E6-0CBC-5BB9-92BD-0C4FE66E1717
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomAttributeResponseBody) GetCustomAttribute() *CustomAttribute {
	return s.CustomAttribute
}

func (s *GetCustomAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomAttributeResponseBody) SetCustomAttribute(v *CustomAttribute) *GetCustomAttributeResponseBody {
	s.CustomAttribute = v
	return s
}

func (s *GetCustomAttributeResponseBody) SetRequestId(v string) *GetCustomAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomAttributeResponseBody) SetSuccess(v bool) *GetCustomAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomAttributeResponseBody) Validate() error {
	if s.CustomAttribute != nil {
		if err := s.CustomAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}
