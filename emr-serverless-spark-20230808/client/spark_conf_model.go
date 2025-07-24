// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkConf interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *SparkConf
	GetKey() *string
	SetValue(v string) *SparkConf
	GetValue() *string
}

type SparkConf struct {
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SparkConf) String() string {
	return dara.Prettify(s)
}

func (s SparkConf) GoString() string {
	return s.String()
}

func (s *SparkConf) GetKey() *string {
	return s.Key
}

func (s *SparkConf) GetValue() *string {
	return s.Value
}

func (s *SparkConf) SetKey(v string) *SparkConf {
	s.Key = &v
	return s
}

func (s *SparkConf) SetValue(v string) *SparkConf {
	s.Value = &v
	return s
}

func (s *SparkConf) Validate() error {
	return dara.Validate(s)
}
