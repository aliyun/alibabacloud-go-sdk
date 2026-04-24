// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProjectToImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AssociateProjectToImageResponseBody
	GetData() *bool
	SetRequestId(v string) *AssociateProjectToImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssociateProjectToImageResponseBody
	GetSuccess() *bool
}

type AssociateProjectToImageResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssociateProjectToImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateProjectToImageResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateProjectToImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *AssociateProjectToImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateProjectToImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssociateProjectToImageResponseBody) SetData(v bool) *AssociateProjectToImageResponseBody {
	s.Data = &v
	return s
}

func (s *AssociateProjectToImageResponseBody) SetRequestId(v string) *AssociateProjectToImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateProjectToImageResponseBody) SetSuccess(v bool) *AssociateProjectToImageResponseBody {
	s.Success = &v
	return s
}

func (s *AssociateProjectToImageResponseBody) Validate() error {
	return dara.Validate(s)
}
