// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFieldDefByUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetFieldDefByUuidResponseBody
	GetRequestId() *string
	SetResult(v []*GetFieldDefByUuidResponseBodyResult) *GetFieldDefByUuidResponseBody
	GetResult() []*GetFieldDefByUuidResponseBodyResult
	SetSuccess(v bool) *GetFieldDefByUuidResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *GetFieldDefByUuidResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetFieldDefByUuidResponseBody
	GetVendorType() *string
}

type GetFieldDefByUuidResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetFieldDefByUuidResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetFieldDefByUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFieldDefByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFieldDefByUuidResponseBody) GetResult() []*GetFieldDefByUuidResponseBodyResult {
	return s.Result
}

func (s *GetFieldDefByUuidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFieldDefByUuidResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetFieldDefByUuidResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetFieldDefByUuidResponseBody) SetRequestId(v string) *GetFieldDefByUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFieldDefByUuidResponseBody) SetResult(v []*GetFieldDefByUuidResponseBodyResult) *GetFieldDefByUuidResponseBody {
	s.Result = v
	return s
}

func (s *GetFieldDefByUuidResponseBody) SetSuccess(v bool) *GetFieldDefByUuidResponseBody {
	s.Success = &v
	return s
}

func (s *GetFieldDefByUuidResponseBody) SetVendorRequestId(v string) *GetFieldDefByUuidResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetFieldDefByUuidResponseBody) SetVendorType(v string) *GetFieldDefByUuidResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetFieldDefByUuidResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFieldDefByUuidResponseBodyResult struct {
	// example:
	//
	// NORMAL
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	Children *string `json:"Children,omitempty" xml:"Children,omitempty"`
	// example:
	//
	// TextareaField
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// textField_laq7xxx
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// example:
	//
	// {}
	Label interface{} `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// {}
	Props interface{} `json:"Props,omitempty" xml:"Props,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFieldDefByUuidResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFieldDefByUuidResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidResponseBodyResult) GetBehavior() *string {
	return s.Behavior
}

func (s *GetFieldDefByUuidResponseBodyResult) GetChildren() *string {
	return s.Children
}

func (s *GetFieldDefByUuidResponseBodyResult) GetComponentName() *string {
	return s.ComponentName
}

func (s *GetFieldDefByUuidResponseBodyResult) GetFieldId() *string {
	return s.FieldId
}

func (s *GetFieldDefByUuidResponseBodyResult) GetLabel() interface{} {
	return s.Label
}

func (s *GetFieldDefByUuidResponseBodyResult) GetProps() interface{} {
	return s.Props
}

func (s *GetFieldDefByUuidResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetFieldDefByUuidResponseBodyResult) SetBehavior(v string) *GetFieldDefByUuidResponseBodyResult {
	s.Behavior = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetChildren(v string) *GetFieldDefByUuidResponseBodyResult {
	s.Children = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetComponentName(v string) *GetFieldDefByUuidResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetFieldId(v string) *GetFieldDefByUuidResponseBodyResult {
	s.FieldId = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetLabel(v interface{}) *GetFieldDefByUuidResponseBodyResult {
	s.Label = v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetProps(v interface{}) *GetFieldDefByUuidResponseBodyResult {
	s.Props = v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetSuccess(v bool) *GetFieldDefByUuidResponseBodyResult {
	s.Success = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
