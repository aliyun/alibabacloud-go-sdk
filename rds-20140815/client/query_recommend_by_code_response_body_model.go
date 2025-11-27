// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecommendByCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryRecommendByCodeResponseBody
	GetData() *string
	SetRequestId(v string) *QueryRecommendByCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRecommendByCodeResponseBody
	GetSuccess() *bool
}

type QueryRecommendByCodeResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {
	//
	//       "SupportOnlineResizeDisk": true,
	//
	//       "DBInstanceName": "rm-bp****",
	//
	//       "maxSupportDiskSizeGB": 6144
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90496720-2319-42A8-87CD-FCE4DF95EBED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecommendByCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecommendByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecommendByCodeResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryRecommendByCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecommendByCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRecommendByCodeResponseBody) SetData(v string) *QueryRecommendByCodeResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRecommendByCodeResponseBody) SetRequestId(v string) *QueryRecommendByCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecommendByCodeResponseBody) SetSuccess(v bool) *QueryRecommendByCodeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRecommendByCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
