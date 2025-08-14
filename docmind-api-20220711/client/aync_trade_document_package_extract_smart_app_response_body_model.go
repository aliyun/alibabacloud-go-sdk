// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAyncTradeDocumentPackageExtractSmartAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompleted(v bool) *AyncTradeDocumentPackageExtractSmartAppResponseBody
	GetCompleted() *bool
	SetCreateTime(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody
	GetCreateTime() *string
	SetData(v interface{}) *AyncTradeDocumentPackageExtractSmartAppResponseBody
	GetData() interface{}
	SetRequestId(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody
	GetRequestId() *string
	SetStatus(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody
	GetStatus() *string
	SetSuccess(v bool) *AyncTradeDocumentPackageExtractSmartAppResponseBody
	GetSuccess() *bool
}

type AyncTradeDocumentPackageExtractSmartAppResponseBody struct {
	Completed  *bool       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	CreateTime *string     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data       interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppResponseBody) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetCompleted(v bool) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Completed = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetCreateTime(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetData(v interface{}) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Data = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetRequestId(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetStatus(v string) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Status = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) SetSuccess(v bool) *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	s.Success = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponseBody) Validate() error {
	return dara.Validate(s)
}
