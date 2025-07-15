// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterVideoResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCasterVideoResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *AddCasterVideoResourceResponseBody
	GetResourceId() *string
}

type AddCasterVideoResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource. This parameter can be used as a request parameter in the API operation that you can call to delete the video source in the production studio or modify the video source in the production studio.
	//
	// example:
	//
	// e5542d98-b08c-46bf-83e9-5b09d08c****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s AddCasterVideoResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterVideoResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterVideoResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterVideoResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddCasterVideoResourceResponseBody) SetRequestId(v string) *AddCasterVideoResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterVideoResourceResponseBody) SetResourceId(v string) *AddCasterVideoResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *AddCasterVideoResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
