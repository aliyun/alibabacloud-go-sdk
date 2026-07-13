// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateResourceCategoryResponseBodyData) *CreateResourceCategoryResponseBody
	GetData() *CreateResourceCategoryResponseBodyData
	SetRequestId(v string) *CreateResourceCategoryResponseBody
	GetRequestId() *string
}

type CreateResourceCategoryResponseBody struct {
	// The returned data.
	Data *CreateResourceCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier of the request.
	//
	// example:
	//
	// AF95C627-D725-5656-B128-B9BCCAF89025
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceCategoryResponseBody) GetData() *CreateResourceCategoryResponseBodyData {
	return s.Data
}

func (s *CreateResourceCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceCategoryResponseBody) SetData(v *CreateResourceCategoryResponseBodyData) *CreateResourceCategoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateResourceCategoryResponseBody) SetRequestId(v string) *CreateResourceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceCategoryResponseBodyData struct {
	// Resource category ID.
	//
	// example:
	//
	// rc-123****7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
}

func (s CreateResourceCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateResourceCategoryResponseBodyData) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *CreateResourceCategoryResponseBodyData) SetResourceCategoryId(v string) *CreateResourceCategoryResponseBodyData {
	s.ResourceCategoryId = &v
	return s
}

func (s *CreateResourceCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
