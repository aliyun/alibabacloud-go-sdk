// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResource(v string) *ModifyCloudResourceResponseBody
	GetCloudResource() *string
	SetRequestId(v string) *ModifyCloudResourceResponseBody
	GetRequestId() *string
}

type ModifyCloudResourceResponseBody struct {
	// The ID of the resource that is added to WAF.
	//
	// example:
	//
	// lb-xxx-80-clb7
	CloudResource *string `json:"CloudResource,omitempty" xml:"CloudResource,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceResponseBody) GetCloudResource() *string {
	return s.CloudResource
}

func (s *ModifyCloudResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudResourceResponseBody) SetCloudResource(v string) *ModifyCloudResourceResponseBody {
	s.CloudResource = &v
	return s
}

func (s *ModifyCloudResourceResponseBody) SetRequestId(v string) *ModifyCloudResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
