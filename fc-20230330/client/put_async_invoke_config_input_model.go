// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAsyncInvokeConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTask(v bool) *PutAsyncInvokeConfigInput
	GetAsyncTask() *bool
	SetDestinationConfig(v *DestinationConfig) *PutAsyncInvokeConfigInput
	GetDestinationConfig() *DestinationConfig
	SetMaxAsyncEventAgeInSeconds(v int64) *PutAsyncInvokeConfigInput
	GetMaxAsyncEventAgeInSeconds() *int64
	SetMaxAsyncRetryAttempts(v int64) *PutAsyncInvokeConfigInput
	GetMaxAsyncRetryAttempts() *int64
}

type PutAsyncInvokeConfigInput struct {
	// example:
	//
	// true
	AsyncTask         *bool              `json:"asyncTask,omitempty" xml:"asyncTask,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// example:
	//
	// 300
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// example:
	//
	// 3
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s PutAsyncInvokeConfigInput) String() string {
	return dara.Prettify(s)
}

func (s PutAsyncInvokeConfigInput) GoString() string {
	return s.String()
}

func (s *PutAsyncInvokeConfigInput) GetAsyncTask() *bool {
	return s.AsyncTask
}

func (s *PutAsyncInvokeConfigInput) GetDestinationConfig() *DestinationConfig {
	return s.DestinationConfig
}

func (s *PutAsyncInvokeConfigInput) GetMaxAsyncEventAgeInSeconds() *int64 {
	return s.MaxAsyncEventAgeInSeconds
}

func (s *PutAsyncInvokeConfigInput) GetMaxAsyncRetryAttempts() *int64 {
	return s.MaxAsyncRetryAttempts
}

func (s *PutAsyncInvokeConfigInput) SetAsyncTask(v bool) *PutAsyncInvokeConfigInput {
	s.AsyncTask = &v
	return s
}

func (s *PutAsyncInvokeConfigInput) SetDestinationConfig(v *DestinationConfig) *PutAsyncInvokeConfigInput {
	s.DestinationConfig = v
	return s
}

func (s *PutAsyncInvokeConfigInput) SetMaxAsyncEventAgeInSeconds(v int64) *PutAsyncInvokeConfigInput {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *PutAsyncInvokeConfigInput) SetMaxAsyncRetryAttempts(v int64) *PutAsyncInvokeConfigInput {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *PutAsyncInvokeConfigInput) Validate() error {
	if s.DestinationConfig != nil {
		if err := s.DestinationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
