// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePrimaryObjectFavoriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TogglePrimaryObjectFavoriteResponseBody
	GetCode() *string
	SetFavoriteCount(v int64) *TogglePrimaryObjectFavoriteResponseBody
	GetFavoriteCount() *int64
	SetMessage(v string) *TogglePrimaryObjectFavoriteResponseBody
	GetMessage() *string
	SetRequestId(v string) *TogglePrimaryObjectFavoriteResponseBody
	GetRequestId() *string
	SetResults(v []*TogglePrimaryObjectFavoriteResponseBodyResults) *TogglePrimaryObjectFavoriteResponseBody
	GetResults() []*TogglePrimaryObjectFavoriteResponseBodyResults
}

type TogglePrimaryObjectFavoriteResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 该用户的关注总数（针对该对象类型）
	//
	// example:
	//
	// 1
	FavoriteCount *int64 `json:"favoriteCount,omitempty" xml:"favoriteCount,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*TogglePrimaryObjectFavoriteResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s TogglePrimaryObjectFavoriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TogglePrimaryObjectFavoriteResponseBody) GoString() string {
	return s.String()
}

func (s *TogglePrimaryObjectFavoriteResponseBody) GetCode() *string {
	return s.Code
}

func (s *TogglePrimaryObjectFavoriteResponseBody) GetFavoriteCount() *int64 {
	return s.FavoriteCount
}

func (s *TogglePrimaryObjectFavoriteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TogglePrimaryObjectFavoriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TogglePrimaryObjectFavoriteResponseBody) GetResults() []*TogglePrimaryObjectFavoriteResponseBodyResults {
	return s.Results
}

func (s *TogglePrimaryObjectFavoriteResponseBody) SetCode(v string) *TogglePrimaryObjectFavoriteResponseBody {
	s.Code = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBody) SetFavoriteCount(v int64) *TogglePrimaryObjectFavoriteResponseBody {
	s.FavoriteCount = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBody) SetMessage(v string) *TogglePrimaryObjectFavoriteResponseBody {
	s.Message = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBody) SetRequestId(v string) *TogglePrimaryObjectFavoriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBody) SetResults(v []*TogglePrimaryObjectFavoriteResponseBodyResults) *TogglePrimaryObjectFavoriteResponseBody {
	s.Results = v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TogglePrimaryObjectFavoriteResponseBodyResults struct {
	// 操作后是否已关注
	//
	// example:
	//
	// true
	IsFavorited *bool `json:"isFavorited,omitempty" xml:"isFavorited,omitempty"`
	// 失败原因（成功时为 null）
	//
	// example:
	//
	// string_value
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 主对象业务ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 操作是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TogglePrimaryObjectFavoriteResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s TogglePrimaryObjectFavoriteResponseBodyResults) GoString() string {
	return s.String()
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) GetIsFavorited() *bool {
	return s.IsFavorited
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) GetObjectId() *string {
	return s.ObjectId
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) SetIsFavorited(v bool) *TogglePrimaryObjectFavoriteResponseBodyResults {
	s.IsFavorited = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) SetMessage(v string) *TogglePrimaryObjectFavoriteResponseBodyResults {
	s.Message = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) SetObjectId(v string) *TogglePrimaryObjectFavoriteResponseBodyResults {
	s.ObjectId = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) SetSuccess(v bool) *TogglePrimaryObjectFavoriteResponseBodyResults {
	s.Success = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
