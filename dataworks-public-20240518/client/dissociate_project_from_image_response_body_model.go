// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateProjectFromImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DissociateProjectFromImageResponseBody
	GetData() *bool
	SetRequestId(v string) *DissociateProjectFromImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DissociateProjectFromImageResponseBody
	GetSuccess() *bool
}

type DissociateProjectFromImageResponseBody struct {
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

func (s DissociateProjectFromImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateProjectFromImageResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DissociateProjectFromImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateProjectFromImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DissociateProjectFromImageResponseBody) SetData(v bool) *DissociateProjectFromImageResponseBody {
	s.Data = &v
	return s
}

func (s *DissociateProjectFromImageResponseBody) SetRequestId(v string) *DissociateProjectFromImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateProjectFromImageResponseBody) SetSuccess(v bool) *DissociateProjectFromImageResponseBody {
	s.Success = &v
	return s
}

func (s *DissociateProjectFromImageResponseBody) Validate() error {
	return dara.Validate(s)
}
