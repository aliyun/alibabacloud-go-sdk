// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICloudPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAICloudPhoneResponseBodyData) *CreateAICloudPhoneResponseBody
	GetData() *CreateAICloudPhoneResponseBodyData
	SetRequestId(v string) *CreateAICloudPhoneResponseBody
	GetRequestId() *string
}

type CreateAICloudPhoneResponseBody struct {
	// The response data object.
	Data *CreateAICloudPhoneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9A51B1DF-96FF-3BCC-B08C-783161D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAICloudPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAICloudPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAICloudPhoneResponseBody) GetData() *CreateAICloudPhoneResponseBodyData {
	return s.Data
}

func (s *CreateAICloudPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAICloudPhoneResponseBody) SetData(v *CreateAICloudPhoneResponseBodyData) *CreateAICloudPhoneResponseBody {
	s.Data = v
	return s
}

func (s *CreateAICloudPhoneResponseBody) SetRequestId(v string) *CreateAICloudPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAICloudPhoneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAICloudPhoneResponseBodyData struct {
	// The order ID.
	//
	// example:
	//
	// 20230930123456
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The list of package IDs. After the payment is successful, instances are created based on these IDs through a callback.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
}

func (s CreateAICloudPhoneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAICloudPhoneResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAICloudPhoneResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateAICloudPhoneResponseBodyData) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *CreateAICloudPhoneResponseBodyData) SetOrderId(v int64) *CreateAICloudPhoneResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateAICloudPhoneResponseBodyData) SetPackageIds(v []*string) *CreateAICloudPhoneResponseBodyData {
	s.PackageIds = v
	return s
}

func (s *CreateAICloudPhoneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
