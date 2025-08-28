// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVirtualNumberRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVirtualNumberRelationResponseBody
	GetCode() *string
	SetData(v string) *QueryVirtualNumberRelationResponseBody
	GetData() *string
	SetRequestId(v string) *QueryVirtualNumberRelationResponseBody
	GetRequestId() *string
}

type QueryVirtualNumberRelationResponseBody struct {
	// The response code.
	//
	// 	- The value 200 indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of associations between virtual numbers and real numbers. The following fields are returned:
	//
	// 	- **relatedNum**: the real number.
	//
	// 	- **createTime**: the time when the number was activated.
	//
	// 	- **pageNo**: the page number.
	//
	// 	- **pageSize**: the number of entries per page.
	//
	// 	- **total**: the total number of entries returned.
	//
	// example:
	//
	// {"data":[{"relatedNum":"1705559****","createTime":"2021-03-26 12:34:08"}],"pageSize":20,"total":1,"pageNo":1}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FAD5988-B483-48A4-B251-6E8470A67371
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVirtualNumberRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVirtualNumberRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVirtualNumberRelationResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryVirtualNumberRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVirtualNumberRelationResponseBody) SetCode(v string) *QueryVirtualNumberRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVirtualNumberRelationResponseBody) SetData(v string) *QueryVirtualNumberRelationResponseBody {
	s.Data = &v
	return s
}

func (s *QueryVirtualNumberRelationResponseBody) SetRequestId(v string) *QueryVirtualNumberRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVirtualNumberRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
