// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlKeywordsFromSqlTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody
	GetCode() *string
	SetData(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody
	GetData() *string
	SetMessage(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody
	GetSuccess() *string
}

type GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The throttling keyword string that was generated based on the SQL statement.
	//
	// example:
	//
	// SELECT~FROM~test~where~name
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30A643F5-D7A6-55F5-AB75-DF501427****
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

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GetData() *string {
	return s.Data
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetCode(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetData(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Data = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetMessage(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetRequestId(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) SetSuccess(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) Validate() error {
	return dara.Validate(s)
}
