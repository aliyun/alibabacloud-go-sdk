// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResourceId(v string) *CreateCloudResourceResponseBody
	GetCloudResourceId() *string
	SetRequestId(v string) *CreateCloudResourceResponseBody
	GetRequestId() *string
}

type CreateCloudResourceResponseBody struct {
	// The ID of the resource that is added to WAF. The ID is automatically generated.
	//
	// example:
	//
	// lb-***
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66A98669-ER12-WE34-23PO-301469*****E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceResponseBody) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *CreateCloudResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudResourceResponseBody) SetCloudResourceId(v string) *CreateCloudResourceResponseBody {
	s.CloudResourceId = &v
	return s
}

func (s *CreateCloudResourceResponseBody) SetRequestId(v string) *CreateCloudResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
