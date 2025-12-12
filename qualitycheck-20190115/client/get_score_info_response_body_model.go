// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScoreInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScoreInfoResponseBody
	GetCode() *string
	SetData(v *GetScoreInfoResponseBodyData) *GetScoreInfoResponseBody
	GetData() *GetScoreInfoResponseBodyData
	SetMessage(v string) *GetScoreInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScoreInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScoreInfoResponseBody
	GetSuccess() *bool
}

type GetScoreInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScoreInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScoreInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScoreInfoResponseBody) GetData() *GetScoreInfoResponseBodyData {
	return s.Data
}

func (s *GetScoreInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScoreInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScoreInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScoreInfoResponseBody) SetCode(v string) *GetScoreInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetData(v *GetScoreInfoResponseBodyData) *GetScoreInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetScoreInfoResponseBody) SetMessage(v string) *GetScoreInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetRequestId(v string) *GetScoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetSuccess(v bool) *GetScoreInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetScoreInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScoreInfoResponseBodyData struct {
	ScorePo []*GetScoreInfoResponseBodyDataScorePo `json:"ScorePo,omitempty" xml:"ScorePo,omitempty" type:"Repeated"`
}

func (s GetScoreInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyData) GetScorePo() []*GetScoreInfoResponseBodyDataScorePo {
	return s.ScorePo
}

func (s *GetScoreInfoResponseBodyData) SetScorePo(v []*GetScoreInfoResponseBodyDataScorePo) *GetScoreInfoResponseBodyData {
	s.ScorePo = v
	return s
}

func (s *GetScoreInfoResponseBodyData) Validate() error {
	if s.ScorePo != nil {
		for _, item := range s.ScorePo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScoreInfoResponseBodyDataScorePo struct {
	// example:
	//
	// 34
	ScoreId    *int32                                         `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreInfos *GetScoreInfoResponseBodyDataScorePoScoreInfos `json:"ScoreInfos,omitempty" xml:"ScoreInfos,omitempty" type:"Struct"`
	ScoreName  *string                                        `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
}

func (s GetScoreInfoResponseBodyDataScorePo) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePo) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePo) GetScoreId() *int32 {
	return s.ScoreId
}

func (s *GetScoreInfoResponseBodyDataScorePo) GetScoreInfos() *GetScoreInfoResponseBodyDataScorePoScoreInfos {
	return s.ScoreInfos
}

func (s *GetScoreInfoResponseBodyDataScorePo) GetScoreName() *string {
	return s.ScoreName
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreId(v int32) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreId = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreInfos(v *GetScoreInfoResponseBodyDataScorePoScoreInfos) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreInfos = v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreName(v string) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreName = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) Validate() error {
	if s.ScoreInfos != nil {
		if err := s.ScoreInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScoreInfoResponseBodyDataScorePoScoreInfos struct {
	ScoreParam []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam `json:"ScoreParam,omitempty" xml:"ScoreParam,omitempty" type:"Repeated"`
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfos) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfos) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfos) GetScoreParam() []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	return s.ScoreParam
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfos) SetScoreParam(v []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) *GetScoreInfoResponseBodyDataScorePoScoreInfos {
	s.ScoreParam = v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfos) Validate() error {
	if s.ScoreParam != nil {
		for _, item := range s.ScoreParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam struct {
	// example:
	//
	// 32
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// example:
	//
	// 3422
	ScoreSubId   *int32  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GetScoreSubId() *int32 {
	return s.ScoreSubId
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GetScoreSubName() *string {
	return s.ScoreSubName
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreNum(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreNum = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreSubId(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreSubId = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreSubName(v string) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreSubName = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreType(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreType = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) Validate() error {
	return dara.Validate(s)
}
