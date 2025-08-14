// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocSendDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchNo(v int64) *PocSendDataRequest
	GetBatchNo() *int64
	SetLang(v string) *PocSendDataRequest
	GetLang() *string
	SetParamsList(v string) *PocSendDataRequest
	GetParamsList() *string
	SetRegId(v string) *PocSendDataRequest
	GetRegId() *string
	SetToken(v string) *PocSendDataRequest
	GetToken() *string
}

type PocSendDataRequest struct {
	// Starting position for batch operations, starting from 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// d0
	BatchNo *int64 `json:"BatchNo,omitempty" xml:"BatchNo,omitempty"`
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Return parameters, in JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// c222460c
	ParamsList *string `json:"ParamsList,omitempty" xml:"ParamsList,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Task token.
	//
	// This parameter is required.
	//
	// example:
	//
	// a471aa421752633438
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s PocSendDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PocSendDataRequest) GoString() string {
	return s.String()
}

func (s *PocSendDataRequest) GetBatchNo() *int64 {
	return s.BatchNo
}

func (s *PocSendDataRequest) GetLang() *string {
	return s.Lang
}

func (s *PocSendDataRequest) GetParamsList() *string {
	return s.ParamsList
}

func (s *PocSendDataRequest) GetRegId() *string {
	return s.RegId
}

func (s *PocSendDataRequest) GetToken() *string {
	return s.Token
}

func (s *PocSendDataRequest) SetBatchNo(v int64) *PocSendDataRequest {
	s.BatchNo = &v
	return s
}

func (s *PocSendDataRequest) SetLang(v string) *PocSendDataRequest {
	s.Lang = &v
	return s
}

func (s *PocSendDataRequest) SetParamsList(v string) *PocSendDataRequest {
	s.ParamsList = &v
	return s
}

func (s *PocSendDataRequest) SetRegId(v string) *PocSendDataRequest {
	s.RegId = &v
	return s
}

func (s *PocSendDataRequest) SetToken(v string) *PocSendDataRequest {
	s.Token = &v
	return s
}

func (s *PocSendDataRequest) Validate() error {
	return dara.Validate(s)
}
