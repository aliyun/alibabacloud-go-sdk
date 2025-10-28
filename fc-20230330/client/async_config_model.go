// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTask(v bool) *AsyncConfig
	GetAsyncTask() *bool
	SetCreatedTime(v string) *AsyncConfig
	GetCreatedTime() *string
	SetDestinationConfig(v *DestinationConfig) *AsyncConfig
	GetDestinationConfig() *DestinationConfig
	SetFunctionArn(v string) *AsyncConfig
	GetFunctionArn() *string
	SetLastModifiedTime(v string) *AsyncConfig
	GetLastModifiedTime() *string
	SetMaxAsyncEventAgeInSeconds(v int64) *AsyncConfig
	GetMaxAsyncEventAgeInSeconds() *int64
	SetMaxAsyncRetryAttempts(v int64) *AsyncConfig
	GetMaxAsyncRetryAttempts() *int64
}

type AsyncConfig struct {
	// example:
	//
	// true
	AsyncTask *bool `json:"asyncTask,omitempty" xml:"asyncTask,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	CreatedTime       *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// example:
	//
	// acs:fc:cn-shanghai:1234/functions/my-func
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// 3600
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// example:
	//
	// 3
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s AsyncConfig) String() string {
	return dara.Prettify(s)
}

func (s AsyncConfig) GoString() string {
	return s.String()
}

func (s *AsyncConfig) GetAsyncTask() *bool {
	return s.AsyncTask
}

func (s *AsyncConfig) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *AsyncConfig) GetDestinationConfig() *DestinationConfig {
	return s.DestinationConfig
}

func (s *AsyncConfig) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *AsyncConfig) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *AsyncConfig) GetMaxAsyncEventAgeInSeconds() *int64 {
	return s.MaxAsyncEventAgeInSeconds
}

func (s *AsyncConfig) GetMaxAsyncRetryAttempts() *int64 {
	return s.MaxAsyncRetryAttempts
}

func (s *AsyncConfig) SetAsyncTask(v bool) *AsyncConfig {
	s.AsyncTask = &v
	return s
}

func (s *AsyncConfig) SetCreatedTime(v string) *AsyncConfig {
	s.CreatedTime = &v
	return s
}

func (s *AsyncConfig) SetDestinationConfig(v *DestinationConfig) *AsyncConfig {
	s.DestinationConfig = v
	return s
}

func (s *AsyncConfig) SetFunctionArn(v string) *AsyncConfig {
	s.FunctionArn = &v
	return s
}

func (s *AsyncConfig) SetLastModifiedTime(v string) *AsyncConfig {
	s.LastModifiedTime = &v
	return s
}

func (s *AsyncConfig) SetMaxAsyncEventAgeInSeconds(v int64) *AsyncConfig {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *AsyncConfig) SetMaxAsyncRetryAttempts(v int64) *AsyncConfig {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *AsyncConfig) Validate() error {
	if s.DestinationConfig != nil {
		if err := s.DestinationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
