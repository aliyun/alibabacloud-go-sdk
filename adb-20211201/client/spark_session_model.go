// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkSession interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v string) *SparkSession
	GetActive() *string
	SetAliyunUid(v int64) *SparkSession
	GetAliyunUid() *int64
	SetSessionId(v int64) *SparkSession
	GetSessionId() *int64
	SetState(v string) *SparkSession
	GetState() *string
}

type SparkSession struct {
	// example:
	//
	// true
	Active *string `json:"Active,omitempty" xml:"Active,omitempty"`
	// example:
	//
	// 11123123
	AliyunUid *int64 `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// example:
	//
	// 15
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// idle
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkSession) String() string {
	return dara.Prettify(s)
}

func (s SparkSession) GoString() string {
	return s.String()
}

func (s *SparkSession) GetActive() *string {
	return s.Active
}

func (s *SparkSession) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *SparkSession) GetSessionId() *int64 {
	return s.SessionId
}

func (s *SparkSession) GetState() *string {
	return s.State
}

func (s *SparkSession) SetActive(v string) *SparkSession {
	s.Active = &v
	return s
}

func (s *SparkSession) SetAliyunUid(v int64) *SparkSession {
	s.AliyunUid = &v
	return s
}

func (s *SparkSession) SetSessionId(v int64) *SparkSession {
	s.SessionId = &v
	return s
}

func (s *SparkSession) SetState(v string) *SparkSession {
	s.State = &v
	return s
}

func (s *SparkSession) Validate() error {
	return dara.Validate(s)
}
