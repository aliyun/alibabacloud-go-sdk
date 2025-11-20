// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DocBlocksQueryResponseBody
	GetRequestId() *string
	SetResult(v *DocBlocksQueryResponseBodyResult) *DocBlocksQueryResponseBody
	GetResult() *DocBlocksQueryResponseBodyResult
	SetSuccess(v bool) *DocBlocksQueryResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *DocBlocksQueryResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DocBlocksQueryResponseBody
	GetVendorType() *string
}

type DocBlocksQueryResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DocBlocksQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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

func (s DocBlocksQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocBlocksQueryResponseBody) GetResult() *DocBlocksQueryResponseBodyResult {
	return s.Result
}

func (s *DocBlocksQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DocBlocksQueryResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DocBlocksQueryResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DocBlocksQueryResponseBody) SetRequestId(v string) *DocBlocksQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocBlocksQueryResponseBody) SetResult(v *DocBlocksQueryResponseBodyResult) *DocBlocksQueryResponseBody {
	s.Result = v
	return s
}

func (s *DocBlocksQueryResponseBody) SetSuccess(v bool) *DocBlocksQueryResponseBody {
	s.Success = &v
	return s
}

func (s *DocBlocksQueryResponseBody) SetVendorRequestId(v string) *DocBlocksQueryResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DocBlocksQueryResponseBody) SetVendorType(v string) *DocBlocksQueryResponseBody {
	s.VendorType = &v
	return s
}

func (s *DocBlocksQueryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DocBlocksQueryResponseBodyResult struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DocBlocksQueryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryResponseBodyResult) GetData() []interface{} {
	return s.Data
}

func (s *DocBlocksQueryResponseBodyResult) SetData(v []interface{}) *DocBlocksQueryResponseBodyResult {
	s.Data = v
	return s
}

func (s *DocBlocksQueryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
