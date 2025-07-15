// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdsFileShareLinksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCdsFileShareLinksResponseBody
	GetCode() *string
	SetData(v []*CdsFileShareLinkModel) *DescribeCdsFileShareLinksResponseBody
	GetData() []*CdsFileShareLinkModel
	SetMessage(v string) *DescribeCdsFileShareLinksResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeCdsFileShareLinksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCdsFileShareLinksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCdsFileShareLinksResponseBody
	GetSuccess() *bool
}

type DescribeCdsFileShareLinksResponseBody struct {
	// The operation result. A value of success indicates that the operation is successful. If the operation failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data information.
	Data []*CdsFileShareLinkModel `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message that is returned. This parameter is not returned if the value of Code is `success`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6mnFXZiT7NdvGNgkInJ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCdsFileShareLinksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdsFileShareLinksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdsFileShareLinksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCdsFileShareLinksResponseBody) GetData() []*CdsFileShareLinkModel {
	return s.Data
}

func (s *DescribeCdsFileShareLinksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCdsFileShareLinksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCdsFileShareLinksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdsFileShareLinksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCdsFileShareLinksResponseBody) SetCode(v string) *DescribeCdsFileShareLinksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCdsFileShareLinksResponseBody) SetData(v []*CdsFileShareLinkModel) *DescribeCdsFileShareLinksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCdsFileShareLinksResponseBody) SetMessage(v string) *DescribeCdsFileShareLinksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCdsFileShareLinksResponseBody) SetNextToken(v string) *DescribeCdsFileShareLinksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCdsFileShareLinksResponseBody) SetRequestId(v string) *DescribeCdsFileShareLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdsFileShareLinksResponseBody) SetSuccess(v bool) *DescribeCdsFileShareLinksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCdsFileShareLinksResponseBody) Validate() error {
	return dara.Validate(s)
}
