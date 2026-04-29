// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteEnterpriseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteEnterpriseTimeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteEnterpriseTimeResponseBody
	GetCode() *string
	SetData(v *CloudDeleteEnterpriseTimeResponseBodyData) *CloudDeleteEnterpriseTimeResponseBody
	GetData() *CloudDeleteEnterpriseTimeResponseBodyData
	SetMessage(v string) *CloudDeleteEnterpriseTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteEnterpriseTimeResponseBody
	GetRequestId() *string
}

type CloudDeleteEnterpriseTimeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteEnterpriseTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8FA46FC5-24CB-53BE-94F6-30CD8D66751C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteEnterpriseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteEnterpriseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteEnterpriseTimeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteEnterpriseTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteEnterpriseTimeResponseBody) GetData() *CloudDeleteEnterpriseTimeResponseBodyData {
	return s.Data
}

func (s *CloudDeleteEnterpriseTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteEnterpriseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteEnterpriseTimeResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteEnterpriseTimeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponseBody) SetCode(v string) *CloudDeleteEnterpriseTimeResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponseBody) SetData(v *CloudDeleteEnterpriseTimeResponseBodyData) *CloudDeleteEnterpriseTimeResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponseBody) SetMessage(v string) *CloudDeleteEnterpriseTimeResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponseBody) SetRequestId(v string) *CloudDeleteEnterpriseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteEnterpriseTimeResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteEnterpriseTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteEnterpriseTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteEnterpriseTimeResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteEnterpriseTimeResponseBodyData) SetResult(v int64) *CloudDeleteEnterpriseTimeResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
