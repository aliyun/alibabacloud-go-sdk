// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlackListStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBlackListStrategyResponseBody
	GetCode() *string
	SetMessage(v string) *QueryBlackListStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBlackListStrategyResponseBody
	GetRequestId() *string
	SetResultObject(v []*QueryBlackListStrategyResponseBodyResultObject) *QueryBlackListStrategyResponseBody
	GetResultObject() []*QueryBlackListStrategyResponseBodyResultObject
}

type QueryBlackListStrategyResponseBody struct {
	// Return code, **200*	- indicates successful API response.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject []*QueryBlackListStrategyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
}

func (s QueryBlackListStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBlackListStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackListStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBlackListStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBlackListStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBlackListStrategyResponseBody) GetResultObject() []*QueryBlackListStrategyResponseBodyResultObject {
	return s.ResultObject
}

func (s *QueryBlackListStrategyResponseBody) SetCode(v string) *QueryBlackListStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBlackListStrategyResponseBody) SetMessage(v string) *QueryBlackListStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBlackListStrategyResponseBody) SetRequestId(v string) *QueryBlackListStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBlackListStrategyResponseBody) SetResultObject(v []*QueryBlackListStrategyResponseBodyResultObject) *QueryBlackListStrategyResponseBody {
	s.ResultObject = v
	return s
}

func (s *QueryBlackListStrategyResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryBlackListStrategyResponseBodyResultObject struct {
	// Blacklist string, separated by **commas**.
	//
	// example:
	//
	// 127.0.0.1,127.0.0.2
	BizContent *string `json:"BizContent,omitempty" xml:"BizContent,omitempty"`
	// List type:
	//
	// - mobile: Phone number blacklist
	//
	// - ip: IP blacklist
	//
	// - identifyNum: ID number blacklist
	//
	// - bankCard: Bank card blacklist
	//
	// example:
	//
	// identifyNum
	BizKey *string `json:"BizKey,omitempty" xml:"BizKey,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1711533786000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 234822
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Product name:
	//
	// - id2meta: ID number two-factor verification
	//
	// - mobile3Meta: Phone number factor verification
	//
	// - bankcardMeta: Bank card factor verification
	//
	// example:
	//
	// id2meta
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Status:
	//
	// - **disabled**: Disabled
	//
	// - **normal**: Enabled
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// User ID.
	//
	// example:
	//
	// 12600512xxxxxxxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryBlackListStrategyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s QueryBlackListStrategyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetBizContent() *string {
	return s.BizContent
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetBizKey() *string {
	return s.BizKey
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetProductName() *string {
	return s.ProductName
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *QueryBlackListStrategyResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetBizContent(v string) *QueryBlackListStrategyResponseBodyResultObject {
	s.BizContent = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetBizKey(v string) *QueryBlackListStrategyResponseBodyResultObject {
	s.BizKey = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetGmtModified(v int64) *QueryBlackListStrategyResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetId(v int64) *QueryBlackListStrategyResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetProductName(v string) *QueryBlackListStrategyResponseBodyResultObject {
	s.ProductName = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetStatus(v string) *QueryBlackListStrategyResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) SetUserId(v int64) *QueryBlackListStrategyResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *QueryBlackListStrategyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
