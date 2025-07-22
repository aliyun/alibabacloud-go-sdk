// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeadLockHistoryResponseBody
	GetCode() *string
	SetData(v string) *GetDeadLockHistoryResponseBody
	GetData() *string
	SetMessage(v string) *GetDeadLockHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeadLockHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDeadLockHistoryResponseBody
	GetSuccess() *string
	SetSynchro(v string) *GetDeadLockHistoryResponseBody
	GetSynchro() *string
}

type GetDeadLockHistoryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//     "total": 2,
	//
	//     "list": [
	//
	//         {
	//
	//             "accountId": "108************",
	//
	//             "textId": "35303d12d52d29ba73bb85fa2d5b****",
	//
	//             "gmtModified": 1732712680000,
	//
	//             "lockTime": 1732687047000,
	//
	//             "gmtCreate": 1732712680000,
	//
	//             "nodeId": "pi-8****************",
	//
	//             "uuid": "pc-8v**************"
	//
	//         },
	//
	//         {
	//
	//             "accountId": "108************",
	//
	//             "textId": "50a24bdcc5fe7e03f92a55ae7574****",
	//
	//             "gmtModified": 1732626448000,
	//
	//             "lockTime": 1722500305000,
	//
	//             "gmtCreate": 1732626448000,
	//
	//             "nodeId": "pi-8****************",
	//
	//             "uuid": "pc-8v**************"
	//
	//         }
	//
	//     ]
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// None
	Synchro *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetDeadLockHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeadLockHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeadLockHistoryResponseBody) GetData() *string {
	return s.Data
}

func (s *GetDeadLockHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeadLockHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeadLockHistoryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDeadLockHistoryResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *GetDeadLockHistoryResponseBody) SetCode(v string) *GetDeadLockHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeadLockHistoryResponseBody) SetData(v string) *GetDeadLockHistoryResponseBody {
	s.Data = &v
	return s
}

func (s *GetDeadLockHistoryResponseBody) SetMessage(v string) *GetDeadLockHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeadLockHistoryResponseBody) SetRequestId(v string) *GetDeadLockHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeadLockHistoryResponseBody) SetSuccess(v string) *GetDeadLockHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeadLockHistoryResponseBody) SetSynchro(v string) *GetDeadLockHistoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetDeadLockHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}
