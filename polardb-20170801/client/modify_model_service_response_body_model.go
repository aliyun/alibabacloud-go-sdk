// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelServiceId(v string) *ModifyModelServiceResponseBody
	GetModelServiceId() *string
	SetRequestId(v string) *ModifyModelServiceResponseBody
	GetRequestId() *string
	SetStatus(v string) *ModifyModelServiceResponseBody
	GetStatus() *string
}

type ModifyModelServiceResponseBody struct {
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A2EE5B4-CC9F-46E1-A747-E43BC9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyModelServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelServiceResponseBody) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *ModifyModelServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyModelServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyModelServiceResponseBody) SetModelServiceId(v string) *ModifyModelServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *ModifyModelServiceResponseBody) SetRequestId(v string) *ModifyModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyModelServiceResponseBody) SetStatus(v string) *ModifyModelServiceResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
