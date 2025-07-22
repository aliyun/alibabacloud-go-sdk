// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeTagResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeTagResponseBodyData) *GetQueryOptimizeTagResponseBody
	GetData() *GetQueryOptimizeTagResponseBodyData
	SetMessage(v string) *GetQueryOptimizeTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeTagResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeTagResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeTagResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned SQL tag data.
	Data *GetQueryOptimizeTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeTagResponseBody) GetData() *GetQueryOptimizeTagResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeTagResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeTagResponseBody) SetCode(v string) *GetQueryOptimizeTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBody) SetData(v *GetQueryOptimizeTagResponseBodyData) *GetQueryOptimizeTagResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeTagResponseBody) SetMessage(v string) *GetQueryOptimizeTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBody) SetRequestId(v string) *GetQueryOptimizeTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBody) SetSuccess(v string) *GetQueryOptimizeTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeTagResponseBodyData struct {
	// The remarks.
	//
	// example:
	//
	// Slow SQL queries of offline synchronization. No optimization is required.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The SQL tags. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// DAS_IN_PLAN,DAS_NOT_IMPORTANT
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetQueryOptimizeTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeTagResponseBodyData) GetComments() *string {
	return s.Comments
}

func (s *GetQueryOptimizeTagResponseBodyData) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeTagResponseBodyData) GetTags() *string {
	return s.Tags
}

func (s *GetQueryOptimizeTagResponseBodyData) SetComments(v string) *GetQueryOptimizeTagResponseBodyData {
	s.Comments = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBodyData) SetSqlId(v string) *GetQueryOptimizeTagResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBodyData) SetTags(v string) *GetQueryOptimizeTagResponseBodyData {
	s.Tags = &v
	return s
}

func (s *GetQueryOptimizeTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
